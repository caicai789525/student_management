<template>
  <div class="add-student-page">
    <h2>新增学生</h2>
    <form @submit.prevent="submitAddStudent">
      <!-- 姓名输入框 -->
      <div class="form-group">
        <label for="name">姓名：</label>
        <input type="text" id="name" v-model="newStudent.name" required>
      </div>
      <!-- 年龄输入框 -->
      <div class="form-group">
        <label for="age">年龄：</label>
        <input type="number" id="age" v-model="newStudent.age" required>
      </div>
      <!-- 年级输入框 -->
      <div class="form-group">
        <label for="grade">年级：</label>
        <input type="text" id="grade" v-model="newStudent.grade" required>
      </div>
      <!-- 专业输入框 -->
      <div class="form-group">
        <label for="major">专业：</label>
        <input type="text" id="major" v-model="newStudent.major" required>
      </div>
      <!-- 学号输入框 -->
      <div class="form-group">
        <label for="student_id">学号：</label>
        <input type="text" id="student_id" v-model="newStudent.student_id" required>
      </div>
      <!-- 班级输入框 -->
      <div class="form-group">
        <label for="class">班级：</label>
        <input type="text" id="class" v-model="newStudent.class" required>
      </div>
      <!-- 性别输入框 -->
      <div class="form-group">
        <label for="gender">性别：</label>
        <input type="text" id="gender" v-model="newStudent.gender" required>
      </div>
      <button type="submit">提交</button>
      <button @click="goBack">返回</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
// 初始化学生信息对象，与模板字段对应
const newStudent = ref({
  name: '',
  age: 0,
  grade: '',
  major: '',
  student_id: '',
  class: '',
  gender: ''
});

const submitAddStudent = async () => {
  const jwtToken = localStorage.getItem('jwtToken');
  try {
    // 发送包含学生信息的 POST 请求
    await axios.post('/students/create', newStudent.value, {
      headers: {
        'Authorization': `Bearer ${jwtToken}`
      }
    });
    router.push('/management');
  } catch (error) {
    console.error('新增学生信息失败', error);
    if (error.response && error.response.status === 401) {
      router.push('/login');
    }
  }
};

const goBack = () => {
    router.push('/management');
};
</script>

<style scoped>
.add-student-page {
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