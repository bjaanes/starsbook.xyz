const projects = [
    {{ range .Projects }}{{ if not .Hidden }}{
        name: '{{ .Name }}',
        shortName: '{{ .ShortName }}',
        img: 'https://starsbook.xyz/{{ .ShortName }}/projectImage.',
    },{{ end }}
    {{ end }}
]

export default projects