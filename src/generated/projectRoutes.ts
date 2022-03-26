import type {RouteRecordRaw} from "vue-router";


const projectRoutes: RouteRecordRaw[] = [
    {
        path: '/ibcfrens',
        name: 'ibcfrens-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/ibcfrens/:id',
        name: 'ibcfrens',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/spunks',
        name: 'spunks-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/spunks/:id',
        name: 'spunks',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/oblitus',
        name: 'oblitus-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/oblitus/:id',
        name: 'oblitus',
        component: () => import('../views/NftView.vue')
    },
    
]

export default projectRoutes