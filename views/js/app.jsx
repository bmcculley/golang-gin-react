class App extends React.Component {
  render() {
    return <Home />;
  }
}

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      jokes: [],
    };

    this.serverRequest = this.serverRequest.bind(this);
  }

  serverRequest() {
    $.get("http://example.com:3000/api/jokes", (res) => {
      this.setState({
        jokes: res,
      });
    });
  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    return (
      <div className="container">
        <br />
        <h2>Jokeish</h2>
        <p>Let's feed you with some funny Jokes!!!</p>
        <div className="row">
          <div className="container">
            {this.state.jokes.map(function (joke, i) {
              return <Joke key={i} joke={joke} />;
            })}
          </div>
        </div>
      </div>
    );
  }
}

class Joke extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      liked: "",
      jokes: [],
    };
    this.like = this.like.bind(this);
    this.serverRequest = this.serverRequest.bind(this);
  }

  like() {
    let joke = this.props.joke;
    this.serverRequest(joke);
  }
  
  serverRequest(joke) {
    $.post(
      "http://example.com:3000/api/jokes/like/" + joke.id,
      { like: 1 },
      (res) => {
        console.log("res... ", res);
        this.setState({ liked: "Liked!", jokes: res });
        this.props.jokes = res;
      }
    );
  }

  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">
            #{this.props.joke.id}{" "}
            <span className="pull-right">{this.state.liked}</span>
          </div>
          <div className="panel-body joke-hld">{this.props.joke.joke}</div>
          <div className="panel-footer">
            {this.props.joke.likes} Likes &nbsp;
            <a onClick={this.like} className="btn btn-default">
              <span className="glyphicon glyphicon-thumbs-up" />
            </a>
          </div>
        </div>
      </div>
    );
  }
}
ReactDOM.render(<App />, document.getElementById("app"));
