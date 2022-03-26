const projects = [
    {{ range .Projects }}{
        name: '{{ .Name }}',
        img: 'https://starsbook.xyz/{{ .ShortName }}/projectImage.'
    },
    {{ end }}
]

export default projects