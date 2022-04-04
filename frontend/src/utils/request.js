import axios from 'axios'
import {Message, MessageBox} from 'element-ui'

const service = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 10000000
})


export default service
