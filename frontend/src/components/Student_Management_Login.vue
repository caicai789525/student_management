<template>
  <div class="login-container">
    <h2>学生信息管理系统 - 登录</h2>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
        <label for="username">用户名:</label>
        <input
            type="text"
            id="username"
            v-model="username"
            placeholder="请输入用户名"
            required
        />
      </div>
      <div class="form-group">
        <label for="password">密码:</label>
        <input
            type="password"
            id="password"
            v-model="password"
            placeholder="请输入密码"
            required
        />
      </div>
      <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
      <button type="submit" :disabled="isLoading">
        {{ isLoading ? '登录中...' : '登录' }}
      </button>
    </form>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import axios from 'axios';
import {useRouter} from 'vue-router';

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const isLoading = ref(false)
const router = useRouter()

const handleLogin = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await axios.post('/auth/login', {
        username: username.value,
        password: password.value
    }, {
        headers: {
            'Content-Type': 'application/json'
        }
    });
    console.log('完整响应对象', response)
    console.log('登录成功', response.data)

    const jwtToken = response.data.token
    if (jwtToken) {
      localStorage.setItem('jwtToken', jwtToken);
      localStorage.setItem('username', username.value);

      // 设置全局 axios 请求头
      axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;
      // 登录成功后跳转到学生管理页面，也就是主页面
      router.push('/management');
    } else {
      errorMessage.value = '未获取到有效的 JWT 令牌';
    }
  } catch (error) {
    console.error('详细错误信息:', error)
    console.error('错误类型:', error.name)
    console.error('错误堆栈:', error.stack)
    if (error.response) {
      errorMessage.value = error.response.data.message || '登录失败，请检查用户名和密码'
    } else if (error.code === 'ECONNABORTED') {
      errorMessage.value = '请求超时，请稍后重试'
    } else if (error.name === 'TypeError') {
      if (error.message.includes('push')) {
        errorMessage.value = '路由跳转出错，请检查路由配置'
      } else if (error.message.includes('axios')) {
        errorMessage.value = 'Axios 请求出错，请检查请求配置'
      } else {
        errorMessage.value = '请求处理出错，请检查代码配置'
      }
    } else if (error.name === 'NetworkError') {
      errorMessage.value = '网络连接异常，请检查网络设置'
    } else {
      errorMessage.value = '网络错误，请稍后重试'
    }
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 100px auto;
  padding: 40px;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  background-color: #ffffff;
}

h2 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: #333;
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 16px;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

button {
  width: 100%;
  padding: 12px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #0056b3;
}

button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.error-message {
  color: #dc3545;
  margin-bottom: 15px;
  text-align: center;
  font-size: 14px;
}
</style>