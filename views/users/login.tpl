<div class="container">
    <h2>Sign In</h2>
    <p>
    <form method="POST" action="/table">
        <div class="form-group">
          <label>Name</label>
          <input type="text" class="form-control" name="username" placeholder=" Your name">
        </div>
        <div class="form-group">
          <label>Email</label>
          <input type="email" class="form-control" name="email" placeholder="Enter email">
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" class="form-control" name="password" placeholder="Password">
        </div>
        <div class="form-group form-check">
          <input type="checkbox" class="form-check-input" id="exampleCheck1">
          <label class="form-check-label" for="exampleCheck1">Policy Agreement</label>
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
    
</div>