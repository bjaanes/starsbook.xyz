import type {RouteRecordRaw} from "vue-router";


const projectRoutes: RouteRecordRaw[] = [
    {{ range .Projects }}{{ if not .Hidden }}{
        path: '/{{ .ShortName }}',
        name: '{{ .ShortName }}-collection',
        component: () => import('../views/CollectionView.vue')
    },
    {
        path: '/{{ .ShortName }}/:id',
        name: '{{ .ShortName }}',
        component: () => import('../views/NftView.vue')
    },{{ end }}
    {{ end }}
]

export default projectRoutes