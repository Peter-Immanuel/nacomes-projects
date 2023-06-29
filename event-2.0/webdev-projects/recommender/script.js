let name1 = document.getElementById("name");
let gender = document.getElementById("gender");
let age = document.getElementById("age");
let genre = document.getElementById("genre");

let button = document.getElementById("submit-btn");

const moviesArray = [
  {
    title: "John Wick",
    imageUrl: "https://images.justwatch.com/poster/304195811/s592/john-wick-chapter-4.webp",
  },
  {
    title: "Avatar",
    imageUrl: "https://images.justwatch.com/poster/304195811/s592/john-wick-chapter-4.webp",
  },
  {
    title: "Spider-Man",
    imageUrl: "https://images.justwatch.com/poster/304195811/s592/john-wick-chapter-4.webp",
  },
];

function createMovieCards() {
  let movieList = document.getElementById("movie-list");

  moviesArray.forEach(function(movie) {
    let titleTag = document.createElement("h4");
    titleTag.innerText = movie.title;
    
    let imageTag = document.createElement("img");
    imageTag.setAttribute("src", movie.imageUrl);
    imageTag.setAttribute("alt", movie.title);


    let movieTag = document.createElement("div");
    movieTag.setAttribute("class", "movie");

    movieTag.appendChild(imageTag);
    movieTag.appendChild(titleTag);

    movieList.appendChild(movieTag);
  });
}

function displayMovies(event) {
  event.preventDefault();

  createMovieCards();
}

button.addEventListener("click", displayMovies);
