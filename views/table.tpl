<div class="container">
    <h1>Welcome {{.username}}</h1>
    <!-- <tr>
        <td>{{.username}}</td>
        <td>{{.email}}</td>
        <td>{{.password}}</td>
    </tr> -->
    
    <tr>
    {{if .username}}
            <table class="table">
                <thead class="thead-light">
                  <tr>
                    <th scope="col">#</th>
                    <th scope="col">Name</th>
                    <th scope="col">Email</th>
                    <th scope="col">Password</th>
                    <th scope="col">Hashed</th>
                    <th scope="col">Matching</th>
                  </tr>
                </thead>
                <tbody>              
                  <tr>
                    <th scope="row">1</th>
                    <td>{{.username}}</td>
                    <td>{{.email}}</td>
                    <td>{{.password}}</td>
                    <td>{{.hashed}}</td>
                    <td>{{.match}}</td>
                  </tr>
                </tbody>
              </table>
    {{end}}
    </tr>
    {{.Decoding}}
    
</div>