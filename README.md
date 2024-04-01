<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/onuraltuntasb/e-commerce">
    <h3>e-commerce</h3>
  </a>

  <h3 align="center">sveltekit-go demo</h3>
  <p align="center">
    <br />
    <a href="https://github.com/onuraltuntasb/e-commerce"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template">View Demo</a>
    ·
    <a href="https://github.com/onuraltuntasb/e-commerce/issues">Report Bug</a>
    ·
    <a href="https://github.com/onuraltuntasb/e-commerce/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project
<br/>
e-commerce-backend is a demo project to see how to develop vanilla golang monolith
backend with postgres that use mostly jsonb data.

e-commerce-frontend is a demo project to see how to develop real world ssr frontend with
sveltekit

The three principles of project:

-   Clean templating for future projects but don't afraid to try new concepts
-   Try to learn sveltekit is ready to production or not
-   Follow golang way (simplestic approach)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

These are some technologies used to develop the project.

-   [![JS][JS.js]][JS-url]
-   [![Sveltekit][Sveltekit]][Sveltekit-url]
-   [![Golang][Golang]][Golang-url]
-   [![Bootstrap][Bootstrap]][Bootstrap-url]
-   [![Vite][Vite]][Vite-url] Vite

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites

-   Node.js v18.0 or upper
-   Npm
-   Go 1.22 or upper
-   Postgres 16

### Installation backend

1. Clone the repo
    ```sh
    git clone https://github.com/onuraltuntasb/e-commerce-backend
    ```
2. Go install dependencies
    ```sh
     go mod tidy
    ```
3. Run e-commerce-backend
    ```js
     go run .\cmd\web\main.go -mode=dev
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Installation frontend

1. Clone the repo
    ```sh
    git clone https://github.com/onuraltuntasb/e-commerce-frontend
    ```
2. Npm install
    ```sh
     npm install
    ```
3. Run e-commerce-frontend
    ```js
     npm run dev
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage
e-commerce is using dynamic product filtering with recursive json merging
![Product list Screen Shot][product-list-screenshot]
<br/>
![Product list1 Screen Shot][product-list1-screenshot]
<br/>
![Account Screen Shot][account-screenshot]


_For more examples, please refer to the [Documentation](https://github.com/onuraltuntasb/feelsafe-ui)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Onur Altuntas - onuraltuntasbusiness@gmail.com

Project Link: [https://github.com/onuraltuntasb/feelsafe-ui](https://github.com/onuraltuntasb/e-commerce)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

Some credits for this project

-   [Choose an Open Source License](https://choosealicense.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/onuraltuntasb/feelsafe-ui.svg?style=for-the-badge
[contributors-url]: https://github.com/onuraltuntasb/feelsafe-ui/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/onuraltuntasb/feelsafe-ui.svg?style=for-the-badge
[forks-url]: https://github.com/onuraltuntasb/feelsafe-ui/network/members
[stars-shield]: https://img.shields.io/github/stars/onuraltuntasb/feelsafe-ui.svg?style=for-the-badge
[stars-url]: https://github.com/onuraltuntasb/feelsafe-ui/stargazers
[issues-shield]: https://img.shields.io/github/issues/onuraltuntasb/feelsafe-ui.svg?style=for-the-badge
[issues-url]: https://github.com/onuraltuntasb/feelsafe-ui/issues
[license-shield]: https://img.shields.io/github/license/onuraltuntasb/feelsafe-ui.svg?style=for-the-badge
[license-url]: https://github.com/onuraltuntasb/feelsafe-ui/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: www.linkedin.com/in/onur-altuntas1
[product-list-screenshot]: readme-assets/e-commerce-product-list.png
[product-list1-screenshot]: readme-assets/e-commerce-product-list1.png
[account-screenshot]: readme-assets/e-commerce-account.png
[postgres-screenshot]: readme-assets/e-commerce-postgres.png
[JS.js]: https://shields.io/badge/JavaScript-F7DF1E?logo=JavaScript&logoColor=000&style=flat-square
[JS-url]: https://www.w3schools.com/js/
[Sveltekit]: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAMAAABF0y+mAAAAXVBMVEVHcEz/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/PgD/////LwD/NQD/OwD/JwD/iXP/8/D/xbr/dFr/zcT/oI//YD//3db/Sx3/u7D/taj/EQBaLP4kAAAADXRSTlMAbrMOnTJFHduJWcft9tsoFwAAASxJREFUKJGNksuigjAMRBFQQByTpi1v/P/PvKG80cXNgkUOaWamjaL/1y0uy+L+G2GuPPtCWQ4YEiEDPC6s0BFpa+eGnoHT7P0FMI3vUBWjPLBUT5TKvZcy5jAa6xjq0LdWPw3jtrKnbrMBeVWkP3lCscIXqJmQq0SVsnXjDjNtTGwUZhJi5tbgucCH/qysFdXku9qK6kF0gu5DxofFHQHpFg1Mr726WZ2onnxzkkP80nfN5AgGycEmdYGNRB8XIlqvJtF4jFhfD0HT+wRvMK3qJyLDHM4XbMemIO8spnjnnCzvghSqikGo7+ac6HChxZyena0MxAeb0V2VrxZdq/v21INPrmbWCF+f0JR8P3Td0Ou2/Sa3dLenVSbRtaaXF+r5hYKqIo7T3+h3/QEhhxqR3YytyAAAAABJRU5ErkJggg==
[Sveltekit-url]: https://kit.svelte.dev/
[Golang]: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAMAAABF0y+mAAAA5FBMVEX////3+Pju7++fpqhbam4kPkUUMzr09PV5hIcQMzsYNz4/U1nS1daIkpQnOT4tOTw7V19HdYJOhJNRjJzf4eJKXGFreHw3S1ExQERCbHdctcxq3/566v955/924f4+XGUnQUh98P9bvdeP7v+78v+S5v6n6/9m0e5z2fVbqsB45f/U9//h8/nr+v/Ex8gjVWLF8f+0ubtaorYAAAA7OThLSknJwsCdlZMaBgDb/v8cMzlZn7J5eHhnYF5xuc5SVlfOm4WHX04cHBydr7TKhmjJjXKTucRtiZC2qqXSu7KZz9/H1NgwLrv1AAABMElEQVR4AVTMVWLDQBADUBkbZl7HzBBmZrj/hbp20e9zRhJSGIZBGsMCYDleEBMCzyWHD1CZbC5fKIqlcqVaq9dr1Uq5JBYbzVy2BXBCu9IRyt0ekfoJSe51y0JH6ahFNLuS1mvrhmlaRKaIZVuSrjia226D9/Sq6wegwvgbAYh8qdeVq00Mho7kjpAICIl/FJElRxwD4/bEB1WYzhjZBsDMpwvTr5c4UKWqAWo2nS6NeIKbTlfrfrUEit1s+6B20+nCMJOJ6T7wtxsWQGMz8SNQ4wF6B8TfMSx/sml8Py0kjtPpDrER+feU5QBgT+fL5XwKo9AmcupJCLne7tTtQYgsp5+U/ny936+nI8f8bfrpkwAIPicthyEJAV7e3l5yuCTl1NXlsEjCwYiTNCIkCQCKFFS1ilhgAwAAAABJRU5ErkJggg==
[Golang-url]: https://go.dev/
[Bootstrap]: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAMAAABF0y+mAAAAnFBMVEWRC/6ICvuHCvuCCvp9Cvh5CfeFAPt3APd1Cfaqa/fo2vzr3fzh0fqgZ/SfU/b////SvPPDo/K9nvHMtfOQTvHPsfhpB/PGq/XfzflzCfXcyfhqB/OROfeWR/XWwvf48P27j/V5DOeDLOqHOeure/SuiPLw5fuQOfF4BO306/yVWe+IJu/68/2da+t2JOr++f5YAOicZ++IRezCo/c66xoZAAABbklEQVR4AVTSVZbDMAxA0RrlhmHigMvMuP+9jeSc0vu9iXnwjjHGMUFJTCn1Zdi3Yb/2QfVB7RPwicz0NgzCMPgqiuGFOkmz/KuiKNM/UKbHKjDaT6UUKABjbFSDeWGjOeOa6lfTdrV7YdkJzvVojEnwOKkQW49dqRmHIKUqg4O31cSZlpDJ6UxzDs10vlhm6Qpw5vXG9qhH6bbHVsAuzQmX6QIItYwCxQn3FmCVHhCVjRrr2oEeT6dH/FFAd2qa8HR2aAYum83SIm5OI80F4rQo8uAa0/ZhsQ+XbsD0rgvA4x5XeIv2DtFGk5tDZPqIC/LYSnDVtcUfl+nFemRqMxOEG5wungbOGLcO0Qg5ZBUgNlcqXAJiXXqkC64zwr879nBorS3z/k/OZT4BISX4lCHMitef8tk4eOUwaztCiyj0+Fo+f6qujxcKdcgm2XflgQxRYNK0mOuzlPP481BpNa+v7P+AkjjlEJIIOWRJAHBVLANkOCX8AAAAAElFTkSuQmCC
[Bootstrap-url]: https://getbootstrap.com/
[Vite]: data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzg3Nzc3Nzc3Nzc2NzU3Ny01Nf/AABEIAB8AIAMBEQACEQEDEQH/xAAZAAACAwEAAAAAAAAAAAAAAAAFBwEEBgP/xAAtEAABAwIEAwcFAQAAAAAAAAABAgMEBREABiExQVGREhQiI3GBsRUkQmGhE//EABsBAAICAwEAAAAAAAAAAAAAAAQFAwYAAgcB/8QALhEAAQMDAgMHAwUAAAAAAAAAAQACAwQRMQUhEkFRImFxobHR8AbB4RMUMoGS/9oADAMBAAIRAxEAPwBnzcyOU+c7GkwrhJ0UlzdPA2thjFQiWMPa5LXVz45Cx7PNSnMKZyLww412dF9sDXlbfFY+opKmgdGxjrcV8d1uoTegLKgEkYVd+qmL2X5BedQlQ8CDck+mE+lCor61kbnmw3NztYfmyNmY2OMkBRCza9UKozCi06xUvxqcc1Skbm1tNMdCfp7GQmUvuOVufRL+E81ZzjTe9Qe9sp86OLqt+SOPTfrjTTp+CTgOD6oSohD+1zCyVKnNMpLaiorccAAA9MKvqvSqmqInjsGRtJJJ/vbPRMNMcyJpackohOmiLGW6dSBoOZ4YpGlae7UatlOMHPcBk+3fZOX9lt0WyLS1R4SqlJF5EvVJO4RuOu/THTdRlbcU8ezWbfPDCSzHtWWoOuhwtUKVmaIC6FVyWR2WXLuMG17DiPb4tiywCLUaUxTC4Ozhj06+6Kpm73bkLvToS67WWICyVR4yQ5KVzPL329zyxXtFpYqCnlrWixlJ4B0Zfb3/AMphVScDN8/dMsAAAAWA2AxGkqnGLEHzVR/rNJWy2B3lvzGFHTxDhfkdv7wwZQ1P7eUE/wATsfBT08ojkBOOa55Qoxo9KAfH3j5/0kG9yCdk3/Xzc8ce11Q2WS0ezG7Dw+eS2q5v1ZNsDCOYCQy//9k=
[Vite-url]: https://vitejs.dev/
[RTK]: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAADVUlEQVR4AdXWA6xlOxQA0D5b37Zt27Zt27Zt27Zt27ZtDE5nJembnFyOcZN1vbvrNoxWj31WeChQw9gsysZswuKMTWA4JYZx2JvX+I+Y/M+r7EQXgWGefBYeIuNXHuIUDucqvqEfVzM+gWGWfHpepC/XMT8thKSO2bmXyFV0DqsKdHAzkZNpJ1CqohPzGP3ZmcBQt34j+vIQY1Ht/4vzK28zydBWoI0H+JeVCFSLaeAKIlsSGOIKzM1vPEYn1WNgdfpzDXUMcffvQORQAoMaOzlf8iZjDU0FziGyOiFpZ1W2ZuoyE7KF43mcKYe0Ag3cwT/MR6CWw/mfyEtsyEosxkyMRV1ux2we0gq08Ai/MSuBsXidWCCjD7/wJleyGRNTM6QVaOZB/mB2Am3cS0z+5HJO5VIe5nP6kvEBhzEpgeqJkxrquIH/WYyQzMudvMAeNOViWpiU1bmMH4i8wVrUUzX5jGxNHccT2YSQ00IXtZQrq545uZS/+YMDaaHigfMSbzMm6xM5lcCQDGUTG/AJfTgyfVf0x4l4nIwLaWVavuUlxqRyMpiIrdmC8QjJwrzLf+xAIORPs1OIXEkPgUZupA8rE6hUgR7uJCbX005IluEHvmJOUqAP/MDrTE7IWYs+3EonocK4z8fvxORHZiEkNRxM5ELq0w8Dv9yOQL7QLh7kfx5mP7opVYGp+IiYvMVEBeVNwrt8xcy9k+RuvmLKMuf80cSkP7uX2YI7eJjIEyxf2KDkFDK2CWlyvcqztJepwEXEnEvLFLwR//E8ExLK9NRGZJyS316fqVCBjfiLyN9sQigwFx/yN+tUmCv54/qCkJbbg3zJlJRby2tzIuukzyFnVl4gchqNZZPDbkSO7P3iWCI7Dea5X8fyvEnkRsamUkwXD/M3K/V+uSC/8D6zVV3vMDFH8xP9uY4JqnR9LXvQj7vpzu/bJ5DxMivSSijQyNTszptEfuFIemhmTGoorHAXe/AbHzMfRbvYWfzLXzzM8ezObhzD7XxFxt/cwZLUMT7n8Rpnsw2rszYH8BR9eYXFCJS8hKzLA/xGLPA3b3E2S9NGfhgf4ldikiV9+ICjKt4NcoW1MTvrszt7sBHzMza1lIprZ3Y2YF8OYReWY8Jc3KjzGABKY6329Mh/ZAAAAABJRU5ErkJggg==
[RTK-url]: https://redux-toolkit.js.org/tutorials/rtk-query
