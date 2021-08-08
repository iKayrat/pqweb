<div class="container">
    <div>
        <p>
        </p>
    </div>
    <div class="container">
        <h4>Log In To Neutron</h4>
    </div>
    <!-- &nbsp;
    {{if .flash.error}}
    <h3>{{.flash.error}}</h3>
    &nbsp;
    {{end}} 
    {{if .Errors}}
    {{range $rec := .Errors}}
    <h3>{{$rec}}</h3>
    {{end}}
    &nbsp;
    {{end}} -->
    <form method="POST">
        <div class="form-column col-md-6">
            <div class="form-row">
                <div class="form-group col-md-4">
                    <a class="btn btn-primary" style="background-color: #3b5998" href="#!" role="button">
                        <i class="fa fa-google"></i> Sign In with Google</a>
                </div>
                <div class="form-group col-md-4">
                    <a class="btn btn-primary" style="background-color: #3b5998" href="#!" role="button">
                        <i class="fa fa-facebook"></i> Sign In with Facebook</a>
                </div>
            </div>

            <!-- Separator line -->

            <div class="form-group col-md-8">
                <div class="form-group">
                    <input type="email" class="form-control" name="Email" placeholder="Username or Email" required>
                    {{if .isExist}} {{.isexist}} {{end}}
                </div>
                <div class="form-group">
                    <input type="password" pattern=".{0}|.{8,}" class="form-control" name="Password"
                        placeholder="Password" required title="Type min 8 characters">
                </div>

                <div class="form-group col-md-12">
                    <div class="fomr-row">
                        <div class="form-group col-md-6">
                            <div class="form-group form-check">
                                <input type="checkbox" class="form-check-input" id="exampleCheck1">
                                <label class="" > Remember me</label>
                            </div>
                        </div>
                        <div class="form-group col-md-6">
                            <a href="/auth/register">Sign Up</a>
                        </div>
                    </div>
                </div>

                <button type="submit" class="btn btn-primary">Create Account</button>

            </div>
        </div>
    </form>
</div>