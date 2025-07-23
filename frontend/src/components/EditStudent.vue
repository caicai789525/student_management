<template>
  <div class="edit-student-container">
    <h2>修改学生信息</h2>
    <form @submit.prevent="submitEdit">
      <div class="form-group">
        <label for="name">姓名:</label>
        <input type="text" id="name" v-model="student.name" required>
      </div>
      <div class="form-group">
        <label for="age">年龄:</label>
        <input type="number" id="age" v-model="student.age" required>
      </div>
      <div class="form-group">
        <label for="grade">年级:</label>
        <input type="text" id="grade" v-model="student.grade" required>
      </div>
      <div class="form-group">
        <label for="major">专业:</label>
        <input type="text" id="major" v-model="student.major" required>
      </div>
      <div class="form-group">
        <label for="student_id">学号:</label>
        <input type="text" id="student_id" v-model="student.student_id" disabled>
      </div>
      <div class="form-group">
        <label for="class">班级:</label>
        <input type="text" id="class" v-model="student.class" required>
      </div>
      <div class="form-group">
        <label for="gender">性别:</label>
        <input type="text" id="gender" v-model="student.gender" required>
      </div>
      <button type="submit" :disabled="isLoading">
        {{ isLoading ? '提交中...' : '提交' }}
      </button>
      <button @click="goBack">返回</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const student = ref({
  name: '',
  age: 0,
  grade: '',
  major: '',
  student_id: '',
  class: '',
  gender: ''
});
const isLoading = ref(false);

// 挂载时获取学生信息
onMounted(async () => {
  const studentId = route.params.id;
  if (!studentId) {
    console.error('未获取到有效的学生 ID，返回管理页面');
    await router.push('/management');
    return;
  }
  const jwtToken = localStorage.getItem('jwtToken');
  if (!jwtToken) {
    await router.push('/login');
    return;
  }
  try {
    isLoading.value = true;
    const response = await axios.get(`/students/${studentId}`, {
      headers: {
        'Authorization': `Bearer ${jwtToken}`
      }
    });
    student.value = response.data;
  } catch (error) {
    console.error('获取学生信息失败', error);
    if (error.response && error.response.status === 401) {
      await router.push('/login');
    }
  } finally {
    isLoading.value = false;
  }
});

// 提交修改信息
const submitEdit = async () => {
  const studentId = route.params.id;
  if (!studentId) {
    console.error('未获取到有效的学生 ID，无法提交修改');
    return;
  }
  const jwtToken = localStorage.getItem('jwtToken');
  if (!jwtToken) {
    await router.push('/login');
    return;
  }
  try {
    isLoading.value = true;
    await axios.put(`/students/${studentId}`, student.value, {
      headers: {
        'Authorization': `Bearer ${jwtToken}`
      }
    });
    await router.push('/management');
  } catch (error) {
    console.error('修改学生信息失败', error);
    if (error.response && error.response.status === 401) {
      await router.push('/login');
    }
  } finally {
    isLoading.value = false;
  }
};

// 返回上一页
const goBack = () => {
  router.go(-1);
};
</script>

<style scoped>
.edit-student-container {
  max-width: 600px;
  margin: 50px auto;
  padding: 30px;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  background-color: #fff;
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
  padding: 12px 24px;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button[type="submit"] {
  background-color: #007bff;
  color: white;
  margin-right: 10px;
}

button[type="submit"]:hover {
  background-color: #0056b3;
}

button:hover {
  opacity: 0.9;
}
</style>