<template>
  <CommonPage>
    <template #action>
      <n-button    type="primary" @click="handleAdd()">
        <i class="i-material-symbols:add mr-4 text-18" />
        创建新规则集
      </n-button>
    </template>
    <MeCrud
      ref="$table"
      v-model:query-items="queryItems"
      :scroll-x="1200"
      :columns="columns"
      :get-data="api.read"
    >

      <MeQueryItem label="状态" :label-width="50">
        <n-select
          v-model:value="queryItems.status"
          clearable
          :options="[
            { label: '所有节点', value: 0 },
            { label: '正在运行', value: 1 },
            { label: '心跳检测失败', value: 2},
            { label: '刚刚注册', value: 3 },
            { label: '取消注册', value: 4 },
            { label: '删除', value: 5 },
          ]"
        />
      </MeQueryItem>
    </MeCrud>

    <MeModal ref="modalRef" width="520px">
      <n-form
          ref="modalFormRef"
          label-placement="left"
          label-align="left"
          :label-width="80"
          :model="modalForm"
          :disabled="modalAction === 'view'"
      >
        <n-form-item
            label="规则集名称"
            path="name"
            :rule="{
            required: true,
            message: '请输入规则集名称',
            trigger: ['input', 'blur'],
          }"
        >
          <n-input v-model:value="modalForm.name" :disabled="modalAction !== 'add'" />
        </n-form-item>
      </n-form>
    </MeModal>
  </CommonPage>
</template>

<script setup>
import { NAvatar, NButton, NSwitch, NTag } from 'naive-ui'
import { formatDateTime } from '@/utils'
import { MeCrud, MeQueryItem, MeModal } from '@/components'
import { useCrud } from '@/composables'
import api from './api'

defineOptions({ name: 'RuleSetList' })

const router = useRouter()
const $table = ref(null)
/** QueryBar筛选参数（可选） */
const queryItems = ref({})

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  { title: 'Id', key: 'id', width: 50, ellipsis: { tooltip: true } },
  { title: '名称', key: 'name', width: 150, ellipsis: { tooltip: true } },
  { title: '流数目', key: 'streamCount', width: 150, ellipsis: { tooltip: true } },
  { title: '规则数目', key: 'ruleCount', width: 150, ellipsis: { tooltip: true } },
  { title: '状态', key: 'statusText', width: 150, ellipsis: { tooltip: true } },
  {
    title: '调度时间',
    key: 'scheduleTime',
    width: 180,
    render(row) {
      return h('span', formatDateTime(row['scheduleTime']))
    },
  },
  { title: '运行节点', key: 'node.name', width: 150, ellipsis: { tooltip: true } },
  { title: '检查状态', key: 'statusCheckText', width: 150, ellipsis: { tooltip: true } },
  {
    title: '最后检查时间',
    key: 'statusCheckTime',
    width: 180,
    render(row) {
      return h('span', formatDateTime(row['statusCheckTime']))
    },
  },

  {
    title: '创建时间',
    key: 'createDate',
    width: 180,
    render(row) {
      return h('span', formatDateTime(row['createTime']))
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 350,
    align: 'right',
    fixed: 'right',
    hideInExcel: true,
    render(row) {
      return [
        h(
            NButton,
            {
              size: 'small',
              type: 'primary',
              onClick: () => handleSchedule(row.id),
            },
            {
              default: () => '调度',
              icon: () => h('i', { class: 'i-fe:tool text-14' }),
            }
        ),
        h(
            NButton,
            {
              size: 'small',
              type: 'primary',
              style: 'margin-left: 6px;',
              onClick: () => handleStop(row.id),
            },
            {
              default: () => '停止',
              icon: () => h('i', { class: 'i-fe:tool text-14' }),
            }
        ),
        h(
            NButton,
            {
              size: 'small',
              type: 'primary',
              style: 'margin-left: 6px;',
              onClick: () =>
                  router.push({ path: `/ruleset/${row.id}`, query: { ruleSetName: row.name } }),
            },
            {
              default: () => '编辑',
              icon: () => h('i', { class: 'i-fe:tool text-14' }),
            }
        ),
        h(
            NButton,
            {
              size: 'small',
              type: 'error',
              style: 'margin-left: 6px;',
              onClick: () => handleDelete(row.id),
            },
            {
              default: () => '删除',
              icon: () => h('i', { class: 'i-material-symbols:delete-outline text-14' }),
            }
        ),
      ]
    },
  },
]

async function handleEnable(row) {
  row.enableLoading = true
  try {
    //await api.enable({ id: row.id, enable: !row.enable })
    row.enableLoading = false
    $message.success('操作成功')
    $table.value?.handleSearch()
  } catch (error) {
    row.enableLoading = false
  }
}

const {
  modalRef,
  modalFormRef,
  modalForm,
  modalAction,
  handleAdd,
  handleDelete,
  handleOpen,
  handleSave,
} = useCrud({
  name: '规则集列表',
  initForm: {},
  doCreate: api.create,
  doDelete: api.delete,
  // doUpdate: api.update,
  refresh: () => $table.value?.handleSearch(),
})


</script>
