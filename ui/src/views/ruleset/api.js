import { request } from '@/utils'

export default {
  create: (data) =>{ request.post('/ruleset', data)},
  read: (params = {}) => request.get('/ruleset', { params }),
  getRuleSet: (id) => request.get('/ruleset/${id}'),
}
