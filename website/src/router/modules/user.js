/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const userRouter = {
  path: '/auth-user',
  component: Layout,
  redirect: '/user/complex-table',
  name: 'Table',
  meta: {
    title: '用户管理',
    icon: 'peoples'
  },
  children: [
    {
      path: 'users',
      component: () => import('@/views/users/auth-users'),
      name: 'users',
      meta: { title: '平台用户' }
    },
    {
      path: 'index',
      component: () => import('@/views/users/rbac/index'),
      name: 'rbac',
      meta: { title: '用户权限' }
    },
    {
      path: 'dynamic-table',
      component: () => import('@/views/table/dynamic-table/index'),
      name: 'DynamicTable',
      meta: { title: 'Dynamic Table' }
    },
    {
      path: 'drag-table',
      component: () => import('@/views/table/drag-table'),
      name: 'DragTable',
      meta: { title: 'Drag Table' }
    },
    {
      path: 'inline-edit-table',
      component: () => import('@/views/table/inline-edit-table'),
      name: 'InlineEditTable',
      meta: { title: 'Inline Edit' }
    }

  ]
}
export default userRouter
