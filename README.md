# DietApp (WIP)

A minimalist, high-performance diet tracking application built to manage my personal nutrition habits. This project combines modern web technologies with a simple, effective interface for tracking daily food intake and nutritional goals.


<div align="center" style="display: flex; flex-wrap: nowrap; justify-content: center; overflow-x: auto; gap: 10px;">
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.03.png" alt="Diet Tracker Screenshot 1" height="300"/>
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.24.png" alt="Diet Tracker Screenshot 2" height="300"/>
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.32.png" alt="Diet Tracker Screenshot 3" height="300"/>
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.37.png" alt="Diet Tracker Screenshot 4" height="300"/>
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.41.png" alt="Diet Tracker Screenshot 5" height="300"/>
  <img src=".github/assets/Screen Shot 2024-12-09 at 13.13.44.png" alt="Diet Tracker Screenshot 6" height="300"/>
</div>

## Motivation

I built this application to solve my own diet tracking needs, focusing on:
- Fast, responsive interface with minimal page refreshes
- Simple data entry that doesn't get in the way
- Easy deployment and maintenance
- Learning some Go along the way
- Spending some time away from bloated JS Frameworks

## Tech Stack

- **Backend**: Go
  - Basic echo web app

- **Frontend**: HTMX + Templ
  - HTMX for dynamic content updates
  - Templ for type-safe HTML templates in Go
  - Extremely nice to work with

- **Database**: Turso
  - SQLite-compatible distributed database
  - Perfect for applications with moderate data requirements

- **Hosting**: AWS
  - Deployed on AWS for reliability and scalability
  - Uses Lambda, API Gateway & Cloudfront
  - CloudWatch for monitoring and logging


## License

MIT License - feel free to use this for your own diet tracking needs!
