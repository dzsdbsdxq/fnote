<template>
  <div class="clearfix">
    <a-upload
      v-model:file-list="fileList"
      :action="serverHost + '/admin-api/files/upload'"
      list-type="picture-card"
      @preview="handlePreview"
      @change="handleChange"
      :before-upload="beforeUpload"
      :headers="{ Authorization: props.authorization }"
      @remove="removeImg"
      name="file"
    >
      <div v-if="fileList && fileList.length < 1">
        <plus-outlined />
        <div style="margin-top: 8px">Upload</div>
      </div>
    </a-upload>
    <a-modal :open="previewVisible" :title="previewTitle" :footer="null" @cancel="handleCancel">
      <loading-outlined v-if="loading"></loading-outlined>
      <img alt="example" style="width: 100%" :src="previewImage" v-else />
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { LoadingOutlined, PlusOutlined } from '@ant-design/icons-vue'
import { ref, watch } from 'vue'
import { message, type UploadChangeParam, type UploadProps } from 'ant-design-vue'

const props = defineProps({
  imageUrl: {
    type: String,
    required: true
  },
  authorization: String
})

watch(
  () => props.imageUrl,
  (newVal) => {
    imgUrl.value = newVal
    fileList.value = []
    if (newVal) {
      fileList.value?.push({
        uid: newVal.split('/').pop() || '',
        name: newVal.split('/').pop() || '',
        status: 'done',
        url: serverHost + newVal,
        thumbUrl: serverHost + newVal
      })
    }
  }
)

const imgUrl = ref<string>(props.imageUrl)
const emit = defineEmits(['update:imageUrl'])
const serverHost = import.meta.env.VITE_API_HOST

function getBase64(file: File) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve(reader.result)
    reader.onerror = (error) => reject(error)
  })
}

const previewVisible = ref(false)
const previewImage = ref('')
const previewTitle = ref('')
const loading = ref<boolean>(false)

const fileList = ref<UploadProps['fileList']>([])
if (imgUrl.value) {
  fileList.value?.push({
    uid: imgUrl.value.split('/').pop() || '',
    name: imgUrl.value.split('/').pop() || '',
    status: 'done',
    url: serverHost + imgUrl.value,
    thumbUrl: serverHost + imgUrl.value
  })
}

const handleCancel = () => {
  previewVisible.value = false
  previewTitle.value = ''
}

const handleChange = async (info: UploadChangeParam) => {
  if (info.file.status === undefined) {
    previewImage.value = ''
    fileList.value = []
  }

  if (info.file.status === 'uploading') {
    loading.value = true
    return
  }

  if (info.file.status === 'done') {
    emit('update:imageUrl', info.file.response.data.url)
    loading.value = false
  }
  if (info.file.status === 'error') {
    previewImage.value = ''
    fileList.value = []
    loading.value = false
    message.error('upload error')
  }
}

const removeImg = () => {
  imgUrl.value = ''
  emit('update:imageUrl', '')
}

// const handlePreview = async (file: UploadProps['fileList'][number]) => { = -!无力吐槽， 官网的写法，ts 无法保证类型安全
const handlePreview = async (file: any) => {
  if (!file.url && !file.preview) {
    file.preview = (await getBase64(file.originFileObj)) as string
  }
  previewImage.value = file.url || file.preview
  previewVisible.value = true
  previewTitle.value = file.name || file.url.substring(file.url.lastIndexOf('/') + 1)
}

// const beforeUpload = (file: UploadProps['fileList'][number]) => { = -!无力吐槽， 官网的写法，ts 无法保证类型安全
const beforeUpload = (file: any) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isJpgOrPng) {
    message.error('You can only upload JPG file!')
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    message.error('Image must smaller than 2MB!')
  }
  return isJpgOrPng && isLt2M
}
</script>
<style scoped>
/* you can make up upload button and sample style by using stylesheets */
.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card .ant-upload-text {
  margin-top: 8px;
  color: #666;
}
</style>
