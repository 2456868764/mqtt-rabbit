<!--------------------------------
 - @Author: Ronnie Zhang
 - @LastEditor: Ronnie Zhang
 - @LastEditTime: 2023/12/05 21:29:56
 - @Email: zclzone@outlook.com
 - Copyright © 2023 Ronnie Zhang(大脸怪) | https://isme.top
 --------------------------------->

<template>
  <CommonPage>
    <template #action>
      <n-button  v-permission="'AddNode'"  type="primary" @click="handleAdd()">
        <i class="i-material-symbols:add mr-4 text-18" />
        创建新节点
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
            label="节点名称"
            path="name"
            :rule="{
            required: true,
            message: '请输入节点名称',
            trigger: ['input', 'blur'],
          }"
        >
          <n-input v-model:value="modalForm.name" :disabled="modalAction !== 'add'" />
        </n-form-item>

        <n-form-item
            label="节点标签"
            path="tag"
            :rule="{
            required: false,
            message: '请输入节点标签',
            trigger: ['input', 'blur'],
          }"
        >
          <n-input v-model:value="modalForm.tag" :disabled="modalAction !== 'add'" />
        </n-form-item>

        <n-form-item
            label="节点IP"
            path="ip"
            :rule="{
            required: true,
            message: '请输入节点IP',
            trigger: ['input', 'blur'],
          }"
        >
          <n-input v-model:value="modalForm.ip" :disabled="modalAction !== 'add'" />
        </n-form-item>

        <n-form-item
            label="节点Port"
            path="portStr"
            :rule="{
            required: true,
            message: '请输入节点Port',
            trigger: ['input', 'blur'],

          }"
        >
          <n-input v-model:value="modalForm.portStr" :disabled="modalAction !== 'add'" />
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

defineOptions({ name: 'NodeMgt' })

const $table = ref(null)
/** QueryBar筛选参数（可选） */
const queryItems = ref({})

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  { title: 'Id', key: 'id', width: 50, ellipsis: { tooltip: true } },
  { title: '节点名称', key: 'name', width: 150, ellipsis: { tooltip: true } },
  { title: '标签', key: 'tag', width: 100, ellipsis: { tooltip: true } },
  { title: 'IP', key: 'ip', width: 150, ellipsis: { tooltip: true } },
  { title: '端口', key: 'port', width: 100, ellipsis: { tooltip: true } },
  { title: '状态', key: 'statusText', width: 150, ellipsis: { tooltip: true } },
  {
    title: '最后心跳时间',
    key: 'heartbeatTime',
    width: 180,
    render(row) {
      return h('span', formatDateTime(row['heartbeatTime']))
    },
  },
  {
    title: '数据源更新时间',
    key: 'lastSourcesTime',
    width: 180,
    render(row) {
      return h('span', formatDateTime(row['LastSourcesTime']))
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
  name: '节点',
  initForm: { port: 8080 },
  doCreate: api.create,
  doDelete: api.delete,
  // doUpdate: api.update,
  refresh: () => $table.value?.handleSearch(),
})


</script>
