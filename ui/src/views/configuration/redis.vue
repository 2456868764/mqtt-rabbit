<template>
  <CommonPage>
  <n-form
      ref="formRef"
      :label-width="80"
      :model="config"
  >

    <n-form-item label="配置内容">
      <n-input
          v-model:value="config.content"
          placeholder="Textarea"
          type="textarea"
          :autosize="{
            minRows: 10,
            maxRows: 20
          }"
      />
    </n-form-item>
    <n-form-item>
      <n-button attr-type="button" @click="handleSave">
        保存
      </n-button>
    </n-form-item>
  </n-form>
  </CommonPage>
</template>

<script setup>
import { NAvatar, NButton, NSwitch, NTag } from 'naive-ui'
import api from './api'

defineOptions({ name: 'RedisMgt' })
const config = ref([])
api.getConfiguration({name: 'redis'}).then(({ data = [] }) => (config.value = data))

async function handleSave() {
  try {
    await api.updateConfiguration(config.value)
    $message.success('操作成功')
  } catch (error) {
    $message.success('操作失败')
  }
}

</script>
