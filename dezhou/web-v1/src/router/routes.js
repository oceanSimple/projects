const homePageRoutes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('../components/login/page.vue'),
    },
    {
        path: '/room',
        name: 'room',
        component: () => import('../components/room/page.vue'),
    },
    {
        path: '/home',
        name: 'home',
        component: () => import('../components/home/page.vue'),
        // redirect: '/home/child1',
        // children: [
        //     {
        //         path: '/home/child1',
        //         name: 'child1',
        //         component: () => import('../components/home/child1.vue'),
        //     },
        //     {
        //         path: '/home/child2',
        //         name: 'child2',
        //         component: () => import('../components/home/child2.vue'),
        //     }
        // ]
    },
]

const defaultPageRoutes = [
    {
        path: '/',
        redirect: '/home',
    }
]


export const route = Array.prototype.concat(homePageRoutes, defaultPageRoutes)