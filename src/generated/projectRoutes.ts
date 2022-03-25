import type {RouteRecordRaw} from "vue-router";


const projectRoutes: RouteRecordRaw[] = [
    {
        path: '/ibcfrens/:id',
        name: 'ibcfrens',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/spunks/:id',
        name: 'spunks',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/oblitus/:id',
        name: 'oblitus',
        component: () => import('../views/NftView.vue')
    },
    
]

export default projectRoutes