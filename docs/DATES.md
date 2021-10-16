# Dates in the Application/Database

## Front-end

Dates should be the user's local time. If someone is looking at their list of meetings, the application sohould display dates/times according to their locale. ISO format and modern software libraries make it trivial to convert a time between timezones.

## Backend/Database

In the database, dates are stored using ISO format and GMT 