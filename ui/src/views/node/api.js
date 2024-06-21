/**********************************
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/05 21:29:51
 * @Email: zclzone@outlook.com
 * Copyright © 2023 Ronnie Zhang(大脸怪) | https://isme.top
 **********************************/

import { request } from '@/utils'

export default {
  create: (data) =>{
      data.port = parseInt(data.portStr, 10)
      request.post('/node/register', data)
  },
  read: (params = {}) => request.get('/node', { params }),
  delete: (id) => request.delete(`/node/${id}`),
}
