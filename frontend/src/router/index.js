import { createRouter, createWebHistory } from 'vue-router';
import StudentLogin from '../components/Student_Management_Login.vue';
import StudentManagementHome from '../components/StudentManagementHome.vue';
import AddStudent from '../components/AddStudent.vue';
import EditStudent from '../components/EditStudent.vue';

const routes = [
  {
    path: '/',
    // 重定向到登录页面
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'StudentLogin',
    component: StudentLogin
  },
  {
    path: '/management',
    name: 'StudentManagementHome',
    component: StudentManagementHome,
    meta: { requiresAuth: true }
  },
  {
    path: '/add-student',
    name: 'AddStudent',
    component: AddStudent,
    meta: { requiresAuth: true }
  },
  {
    path: '/edit-student/:id',
    name: 'EditStudent',
    component: EditStudent,
    meta: { requiresAuth: true }
  },
  {
    path:'/delete-student/:studentId',
    name:'DeleteStudent',
    meta: { requiresAuth: true }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const jwtToken = localStorage.getItem('jwtToken');
    if (!jwtToken) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;