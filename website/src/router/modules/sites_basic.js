/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const sitesRouter = {
  path: '/',
  component: Layout,
  redirect: '/sites',
  name: 'Table',
  meta: {
    title: 'resources',
    icon: 'peoples'
  },
  children: [
    {
      path: '/',
      component: () => import('@/views/sites/index'),
      name: 'sites'
      // redirect: '/sites',
      //  meta: { title: 'z-gour.com', icon: 'dashboard', affix: true }

    },
    {
      path: '/search',
      component: () => import('@/views/sites/search/index'),
      name: 'search'
      // meta: { title: 'resources' }
    },
    {
      path: '/about',
      component: () => import('@/views/sites/about/index'),
      name: 'about',
      meta: { title: 'resources' }
    }

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
export default sitesRouter
