import { request } from '@/utils'

export default {
  create: (data) =>{
      data.port = parseInt(data.portStr, 10)
      request.post('/node/register', data)
  },
  read: (params = {}) => request.get('/node', { params }),
  delete: (id) => request.delete(`/node/${id}`),
}
