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
    {
        path: '/alphacentaurians',
        name: 'alphacentaurians-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/alphacentaurians/:id',
        name: 'alphacentaurians',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/114shut',
        name: '114shut-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/114shut/:id',
        name: '114shut',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/acre',
        name: 'acre-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/acre/:id',
        name: 'acre',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/starty',
        name: 'starty-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/starty/:id',
        name: 'starty',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/trooprs',
        name: 'trooprs-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/trooprs/:id',
        name: 'trooprs',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/hodlavatars',
        name: 'hodlavatars-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/hodlavatars/:id',
        name: 'hodlavatars',
        component: () => import('../views/NftView.vue')
    },
    {
        path: '/sunnysidereapers',
        name: 'sunnysidereapers-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/sunnysidereapers/:id',
        name: 'sunnysidereapers',
        component: () => import('../views/NftView.vue')
    },
    
]

export default projectRoutes