<template>
  <CommonPage back>
    <template #title-suffix>
      <n-tag class="ml-12" type="warning">{{ route.query.ruleSetName }}</n-tag>
    </template>
    <template #action>
      <div class="flex items-center">
        <n-button type="primary" @click="handleStreamAdd()">
          <i  class="i-line-md:confirm-circle mr-4 text-18" />
          添加流
        </n-button>
        <n-button
            class="ml-12"
            type="primary"
            @click="handleRuleAdd()"
        >
          <i class="i-line-md:confirm-circle mr-4 text-18" />
          添加规则
        </n-button>
      </div>
    </template>

    <n-card title="流">
        <MeCrud
            ref="$table"
            v-model:query-items="queryItems"
            :scroll-x="1200"
            :columns="columns"
            :get-data="api.read"
        >
        </MeCrud>
    </n-card>

    <n-card title="规则">
      <MeCrud
          ref="$table"
          v-model:query-items="queryItems"
          :scroll-x="1200"
          :columns="columns"
          :get-data="api.read"
      >
      </MeCrud>
    </n-card>

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

defineOptions({ name: 'RuleSetEdit' })
const route = useRoute()
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
    width: 100,
    align: 'right',
    fixed: 'right',
    hideInExcel: true,
    render(row) {
      return [
        h(
            NButton,
            {
              size: 'small',
              type: 'error',
              style: 'margin-left: 12px;',
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