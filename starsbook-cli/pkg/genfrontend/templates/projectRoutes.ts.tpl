import type {RouteRecordRaw} from "vue-router";


const projectRoutes: RouteRecordRaw[] = [
    {{ range .Projects }}{
        path: '/{{ .ShortName }}',
        name: '{{ .ShortName }}-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/{{ .ShortName }}/:id',
        name: '{{ .ShortName }}',
        component: () => import('../views/NftView.vue')
    },
    {{ end }}
]

export default projectRoutes