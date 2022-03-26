const projects = [
    {{ range .Projects }}{
        name: '{{ .Name }}',
        shortName: '{{ .ShortName }}',
        img: 'https://starsbook.xyz/{{ .ShortName }}/projectImage.'
    },
    {{ end }}
]

export default projects