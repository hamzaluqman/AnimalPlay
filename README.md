# AnimalPlay

Feeling sad? Need to hug a goat or a bunny? This web application will allow a user to find animals available to play with as a form of enjoyment for the animal and the user.

Animals will belong to a veterinary school.

Web app will consist of a Go backend, React.js frontend, MySQL database and utilize a hexagonal architecture (Without ports). Ports will not be used since Go has implicit interface satisfaction and
it is prefered to define interfaces at the consumer level instead of the production level.

An email service will be created to send a booking confirmation email to the user.

Email and web application will be separate containerized services .
