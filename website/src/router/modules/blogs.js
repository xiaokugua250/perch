/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const blogsRouter = {
  path: '/blogs',
  component: Layout,
  redirect: 'noRedirect',
  name: 'Table',
  meta: {
    title: 'resources',
    icon: 'peoples'
  },
  children: [
/*
    {
      path: 'show',
      component: () => import('@/views/resources/blogs/show'),
      name: 'blogs_show',
    //  meta: { title: 'blogs_show' }
    },
    {
      path: 'list',
      component: () => import('@/views/resources/blogs/list'),
      name: 'blogs_list',
      meta: { title: 'resources' }
    },
    {
      path: '/',
      component: () => import('@/views/resources/blogs/index'),
      name: 'resources'
      // meta: { title: 'resources' }
    }*/

  ]
}
export default blogsRouter
