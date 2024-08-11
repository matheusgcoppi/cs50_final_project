Finance App
Video Demo: https://youtu.be/SfnGFnctMMA
Description:
The Finance App is designed to help individuals manage and track their finances more effectively. Users can input their incomes and expenses, and the app displays this data visually in the form of charts, making it easier to understand financial trends over the week and year. This tool is particularly useful for anyone who wants to gain a clearer picture of their financial situation and make informed decisions based on that data.

The application is divided into two parts: a backend API built with Go, and a frontend built with Vue.js. The backend manages the business logic and database interactions, while the frontend handles user interaction and visualization.

Key Features:
Income and Expense Tracking: Easily input and track your financial transactions.
Weekly and Yearly Charts: Visualize your income and expenses over the week and year to see how much you are earning and spending.
User Authentication: Securely manage your account with JWT-based authentication.
Password Reset via Email: Users can reset their passwords by requesting a password reset email. The system will send an email with instructions to securely reset the password.
Technical Details:
Technologies Used:
Backend: Go, Echo framework, PostgreSQL
Frontend: Vue.js, Chart.js for visualizations
Database: PostgreSQL
Email Service: Configured SMTP service to send password reset emails.
Project Structure:
The project is split into two main sections: frontend and backend.
The backend is configured to handle requests via a RESTful API, and it communicates with the PostgreSQL database.
The frontend is a Vue.js application that consumes the API and provides an interface for users to interact with.
The backend includes an email service to send password reset emails.
Communication:
The frontend communicates with the backend via API calls. This allows the frontend to send user data (like incomes and expenses) to the backend and retrieve processed data for visualization.
The password reset functionality uses the email service configured in the backend to send secure password reset emails to users.
Challenges:
One of the main challenges faced during development was working with charts to ensure that the data is displayed correctly and efficiently. This was overcome by leveraging the Chart.js library, which provided robust charting capabilities integrated into the Vue.js frontend.
Implementing the password reset functionality involved configuring a secure SMTP service and ensuring that sensitive data is protected during the reset process.
Setup Instructions:
Backend:
Clone the repository:
bash
Copy code
git clone https://github.com/matheusgcoppi/cs50_final_project/
cd finance-app/backend
Configure the .env file:
Provide your database credentials, JWT secret, email SMTP settings, and other necessary configurations.
Example:
makefile
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
SMTP_HOST=smtp.your-email-provider.com
SMTP_PORT=587
SMTP_USER=youremail@domain.com
SMTP_PASSWORD=your-email-password
Run the backend:
The project includes a Makefile for convenience. Simply run:
bash
Copy code
make run
This will start the Go server with the Echo framework, connecting to your PostgreSQL database and enabling the email service for password resets.
Frontend:
Navigate to the frontend directory:
bash
Copy code
cd ../frontend
Install dependencies:
bash
Copy code
npm install
Run the frontend server:
bash
Copy code
npm run serve
This will start the Vue.js development server, which will run on a different port from the backend.
Environment Variables:
Ensure that the .env file in the backend is properly configured with your database credentials, JWT secret, and SMTP settings.
Files and Directories:
Backend:
cmd: Contains the main Go file that starts the application.
database: Configures the connection to the PostgreSQL database.
mail: Contains email setup and configurations, including the SMTP settings for sending password reset emails.
middleware: Holds middleware functions for the Echo framework, such as logging and authentication.
repository: Contains the SQL queries and database access functions.
service: Implements business logic and core functionalities, including handling password reset requests.
Frontend:
components: Reusable Vue components (e.g., forms, buttons, input fields).
composables: Contains Vue.js composables (functions) that handle shared logic and state management.
views: Defines the different pages/views of the application (e.g., Finance Dashboard, Income Entry, Expense Entry).
Conclusion:
The Finance App is a practical tool designed to help users take control of their financial data by making it easier to track and visualize their incomes and expenses. The combination of Go for the backend, Vue.js for the frontend, and PostgreSQL as the database makes for a robust and scalable application that can be easily deployed and used by individuals looking to better manage their finances. Additionally, the password reset functionality ensures that users can securely regain access to their accounts if they forget their passwords.

