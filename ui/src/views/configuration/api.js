
import { request } from '@/utils'

export default {
  getConfiguration: (params = {}) => request.get('/configuration', { params }),
  updateConfiguration: (data) => request.put(`/configuration`, data),
}
