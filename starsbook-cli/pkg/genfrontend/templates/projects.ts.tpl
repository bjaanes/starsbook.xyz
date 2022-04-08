const projects = [
    {{ range .Projects }}{{ if not .Hidden }}{
        name: '{{ .Name }}',
        shortName: '{{ .ShortName }}',
        img: 'https://starsbook-assets.storage.googleapis.com/{{ .ShortName }}/projectImage.',
    },{{ end }}
    {{ end }}
]

export default projects