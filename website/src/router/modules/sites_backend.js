/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const sitesBackendRouter = {
  path: '/backend',
  component: Layout,
 // redirect: 'backend',
  name: 'backend',
  meta: {
    title: 'resources',
    icon: 'peoples'
  },
  children: [
   {
      path: 'user',
      component: () => import('@/views/sites/backend/user/index'),
      name: 'user',
      // meta: { title: 'resources' }
    },
    {
      path: 'monitor',
      component: () => import('@/views/sites/backend/monitor/monitor'),
      name: 'monitor',
      // meta: { title: 'resources' }
    },
    {
      path: 'k8scloud',
     // redirect:"k8scloud/a",
     // component: Layout,
       component: () => import('@/views/sites/backend/k8scloud/k8scloud'),
       name: 'k8scloud',
      children:[
/*
        {
          path: '',
          component: () => import('@/views/sites/backend/k8scloud/k8scloud'),
       //   name: 'basic',
        },*/
           {
             path: 'submit',
             component: () => import('@/views/sites/backend/k8scloud/k8scloud_submit'),
         //    name: 'submit',
           }
      ]
      // meta: { title: 'resources' }
    },
    {
      path: 'filesystem',
      component: () => import('@/views/sites/backend/filesystem/index'),
      name: 'user',
      // meta: { title: 'resources' }
    },
    {
      path: 'resources',
      component: () => import('@/views/sites/backend/resources/index'),
      name: 'user',
      // meta: { title: 'resources' }
    },
    {
      path: 'spider',
      component: () => import('@/views/sites/backend/spider/index'),
      name: 'user',
      // meta: { title: 'resources' }
    },
    {
      path: '/',
      component: () => import('@/views/sites/backend/index'),
      name: 'backend'
      // redirect: '/sites',
      //  meta: { title: 'z-gour.com', icon: 'dashboard', affix: true }

    },


    /*
     {
    path: '/',
    component: Layout,
    redirect: '/site',
    children: [
      {
        path: '/',
        component: () => import('@/views/sites/index'),
        name: 'sites',
        meta: { title: 'z-gour.com', icon: 'dashboard', affix: true }
      }
    ]
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
*/
  ]
}
export default sitesBackendRouter
