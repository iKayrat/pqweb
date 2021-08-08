<div class="container">
    <p>
        <h3>{{.Error}}</h3>
{{if .firstname}}
    <h2>Welcome {{.firstname}}</h2>
    lastname: {{.lastname}}
    username: {{.username}}
    email: {{.email}}
    password: {{.password}}
    {{end}}
</div>