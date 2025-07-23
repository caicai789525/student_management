<template>
  <div class="management-page">
    <!-- 顶部导航栏 -->
    <header class="header">
      <div class="header-title">学生信息管理系统</div>
      <div class="user-section">
        <img src="@/assets/logo.png" alt="用户头像" class="user-avatar" />
        <span class="user-name">{{ username }}</span>
        <div class="dropdown">
          <button @click="showDropdown =!showDropdown" class="dropdown-toggle">▼</button>
          <div v-if="showDropdown" class="dropdown-menu">
            <button @click="logout" class="logout-btn">退出登录</button>
          </div>
        </div>
      </div>
    </header>
    <div class="main-container">
      <!-- 左侧侧边栏 -->
      <aside class="sidebar">
        <button @click="goToStudentManagement" class="nav-btn">
          <i class="folder-icon"></i> 学生管理
        </button>
      </aside>
      <!-- 主要内容区 -->
      <main class="main-content">
        <button @click="goToAddStudent" class="add-btn">新增</button>
        <table class="student-table">
          <thead>
          <tr>
            <th>学号</th>
            <th>姓名</th>
            <th>性别</th>
            <th>年龄</th>
            <th>班级</th>
            <th>专业</th>
            <th>操作</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="student in students" :key="student.student_id">
            <td>{{ student.student_id }}</td>
            <td>{{ student.name }}</td>
            <td>{{ student.gender }}</td>
            <td>{{ student.age }}</td>
            <td>{{ student.class }}</td>
            <td>{{ student.major }}</td>
            <td>
              <button @click="goToEditStudent(student.student_id)" class="edit-btn">编辑</button>
              <button @click="deleteStudent(student.student_id)" class="delete-btn">删除</button>
            </td>
          </tr>
          </tbody>
        </table>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const students = ref([]);
const showDropdown = ref(false);
const username = ref(localStorage.getItem('username') || '');

onMounted(async () => {
  const jwtToken = localStorage.getItem('jwtToken');
  if (!jwtToken) {
    router.push('/login');
    return;
  }
  try {
    const response = await axios.get('/students', {
      headers: {
        'Authorization': `Bearer ${jwtToken}`
      }
    });
    students.value = response.data;
  } catch (error) {
    console.error('获取学生信息失败', error);
    if (error.response && error.response.status === 401) {
      await router.push('/login');
    }
  }
});

const logout = () => {
  localStorage.removeItem('jwtToken');
  localStorage.removeItem('username');
  delete axios.defaults.headers.common['Authorization'];
  router.push('/login');
};

const goToAddStudent = () => {
  router.push('/add-student');
};

const goToEditStudent = (studentId) => {
  router.push(`/edit-student/${studentId}`);
};

const deleteStudent = async (studentId) => {
  const jwtToken = localStorage.getItem('jwtToken');
  try {
    await axios.delete(`/students/${studentId}`, {
      headers: {
        'Authorization': `Bearer ${jwtToken}`
      }
    });
    students.value = students.value.filter(student => student.student_id!== studentId);
  } catch (error) {
    console.error('删除学生信息失败', error);
    if (error.response && error.response.status === 401) {
      await router.push('/login');
    }
  }
};

const goToStudentManagement = () => {
  // 可添加跳转逻辑，当前保持在本页面
};
</script>

<style scoped>
.management-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #007bff;
  padding: 15px 30px;
}

.header-title {
  color: white;
  font-size: 1.5rem;
  font-weight: bold;
}

.user-section {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-name {
  color: white;
  font-size: 1rem;
}

.dropdown {
  position: relative;
}

.dropdown-toggle {
  background: none;
  border: none;
  color: white;
  font-size: 1.2rem;
  cursor: pointer;
}

.dropdown-menu {
  position: absolute;
  right: 0;
  top: 100%;
  background-color: white;
  border-radius: 4px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 8px 0;
  min-width: 120px;
}

.logout-btn {
  width: 100%;
  padding: 8px 16px;
  background: none;
  border: none;
  text-align: left;
  cursor: pointer;
}

.logout-btn:hover {
  background-color: #f0f0f0;
}

.main-container {
  display: flex;
  flex: 1;
}

.sidebar {
  background-color: #004085;
  width: 200px;
  padding: 20px 0;
  height: 100%; /* 确保侧边栏高度占满父容器 */
}

.nav-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  padding: 12px 20px;
  background: none;
  border: none;
  color: white;
  font-size: 1rem;
  cursor: pointer;
}

.nav-btn:hover {
  background-color: #0056b3;
}

.folder-icon {
  display: inline-block;
  width: 20px;
  height: 20px;
  background: url('@/assets/logo.png') no-repeat center/contain;
}

.main-content {
  flex: 1;
  background-color: white;
  padding: 30px;
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow-y: auto; /* 当内容超出时显示滚动条 */
}

.add-btn {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 10px 20px;
  font-size: 1rem;
  cursor: pointer;
  margin-bottom: 20px;
  align-self: flex-start;
}

.add-btn:hover {
  background-color: #0056b3;
}

.student-table {
  width: 100%;
  border-collapse: collapse;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.student-table th,
.student-table td {
  padding: 12px 16px;
  text-align: left;
  border-bottom: 1px solid #dee2e6;
}

.student-table th {
  background-color: #f8f9fa;
  font-weight: 600;
}

.edit-btn {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 6px 12px;
  cursor: pointer;
  margin-right: 8px;
}

.edit-btn:hover {
  background-color: #0056b3;
}

.delete-btn {
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 6px 12px;
  cursor: pointer;
}

.delete-btn:hover {
  background-color: #c82333;
}
</style>