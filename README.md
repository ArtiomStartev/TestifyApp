## 🧪 TestifyApp

**TestifyApp** is a comprehensive full-stack web application built with **Go** and **Fiber**, designed to streamline the process of **creating, taking, and managing tests**. It features **JWT-based authentication** and **role-based access control**, ensuring a secure, personalized experience for both regular users and admins.

- **For Users**:  
  Sign up, log in, and begin testing right away. Get **instant feedback** on your performance with detailed results including score, pass/fail status, and time taken. Track improvement across multiple attempts via a dedicated results dashboard.

- **For Admins**:  
  Elevated privileges grant access to an **All Users** overview, enabling you to **monitor every user** in the system and inspect their results with a single click. Quickly identify high performers, troubleshoot problem areas, and maintain full oversight of the testing platform.

Whether you’re **preparing for an exam**, **onboarding new hires**, or **just testing knowledge**, **TestifyApp** offers a clean, user-friendly interface to manage the entire lifecycle of test administration and result analysis.


---

## 🔐 Login Page

The **Login Page** allows registered users to access their accounts securely. It features a minimalistic and centered form design with a smooth user experience.

### ✨ Features

- Clean and user-friendly login form
- Email and password input fields with proper placeholders
- Prominent **Login** button
- Link to sign up for new users (`Sign up`)
- Error handling for incorrect credentials

### ✅ User Authentication & Authorization

- Sign-up and login flows secured with **JWT tokens**
- Passwords hashed using `bcrypt` for strong security
- JWT stored in **HTTP-only cookies** to protect against XSS attacks and ensure safe session management

This page acts as the secure entry point to the platform, ensuring that only verified users gain access.

<img width="1790" alt="Screenshot 2025-04-06 at 09 31 23" src="https://github.com/user-attachments/assets/a27715f5-475e-4595-8fc3-4cb4fa493d7c" />


## 🏠 Home Page

After logging in, users are redirected to the **Home Page**, where they are welcomed with a personalized message and a clear **"Start Test"** button.

### ✨ Features

- Displays a personalized greeting with the user's email
- Prominent **Start Test** button to launch the test flow
- Minimal and clean design for a distraction-free experience
- Top navigation bar with role-based access:
  - **Home**
  - **Your Results**
  - **Logout**
  - ✅ **Admin Only:** `All Users` tab is visible and accessible only for admin accounts

### 👤 Role-Based Navigation

| Role   | Visible Tabs                    |
|--------|----------------------------------|
| User   | Home, Your Results, Logout       |
| Admin  | Home, Your Results, All Users, Logout |

This page serves as a hub, guiding users (and admins) into their next action with clarity and simplicity.

<img width="1790" alt="Screenshot 2025-04-06 at 09 28 06" src="https://github.com/user-attachments/assets/874706a6-4b5b-4e5b-8f07-58272c0957d6" />


## 🏁 Test System

When users click **Start Test**, they begin a structured, time-tracked assessment with a mix of **single-choice** and **multiple-choice** questions.

### ✨ Features

- **Timed Tests**  
  A live timer in the corner shows how long the user has spent, making each attempt measurable.

- **Dynamic Question Navigation**  
  Intuitive **Back** and **Next** buttons let users review previous questions and progress through the test at their own pace.

- **Multiple Question Types**  
  - Single-choice (radio) questions for scenarios with one correct answer  
  - Multiple-choice (checkbox) questions for scenarios with multiple correct answers

- **Instant Submission & Feedback**  
  Once the user hits **Finish Test**, results are:
  - Immediately **evaluated** on the server  
  - **Scored** and displayed with a clear pass/fail indicator and the number of correct answers

- **Progress Tracking**  
  Every test attempt is recorded, allowing users to see how many questions they got right, how much time they took, and whether they passed.

This interactive and user-friendly test system offers a seamless experience, ensuring both casual learners and advanced users stay engaged and informed during their assessments.

<img width="1788" alt="Screenshot 2025-04-06 at 09 41 38" src="https://github.com/user-attachments/assets/f68628b1-d7c8-485e-93d1-2f7058659b96" />

<img width="1787" alt="Screenshot 2025-04-06 at 09 44 08" src="https://github.com/user-attachments/assets/b92b921d-f48b-4e54-82c7-8c14668609e6" />

<img width="1788" alt="Screenshot 2025-04-06 at 09 44 48" src="https://github.com/user-attachments/assets/249de046-53f7-42d4-b85d-6ee4c0c224e7" />


## 📊 User Dashboard

Users can review their **personal test history** and monitor progress over time via the **Your Results** page.

### ✨ Features

- **Consolidated View**  
  Displays all test attempts in a single table, with each row showing key data at a glance

- **Detailed Metrics**  
  - **Attempt #**: Distinguishes each consecutive test session
  - **Score**: Number of correct answers
  - **Result**: Clear pass/fail status  
  - **Duration**: How long the user spent on the test in seconds

- **Progress Tracking**  
  Users can quickly identify patterns (e.g., improved performance over multiple attempts)

- **Intuitive Layout**  
  The table is easy to read, and results are color-coded (green for pass, red for fail)

This dashboard empowers users with immediate feedback on their performance, encouraging them to identify areas for improvement and keep track of their success over time.

<img width="1787" alt="Screenshot 2025-04-06 at 09 51 00" src="https://github.com/user-attachments/assets/ef7602b0-7c1f-4886-81c6-769a72ed6068" />


## 🔐 Admin Dashboard

For users flagged as `is_admin = true`, the **All Users** page provides complete visibility and control over the platform's user base.

### ✨ Features

- **Role-Based Access**
  - Only admins can access this page
  - Admin status is set via a database flag (`is_admin = true`)

- **Comprehensive User List**
  - Displays email and user ID for each account
  - Clean, card-based layout makes it easy to scan through the list

- **Expandable User Results**
  - Click or tap a user to reveal all of their test attempts
  - Each attempt shows score, pass/fail status, and duration
  - Collapsible sections keep the interface clutter-free

### 👀 Monitoring & Insights
- Track overall performance of the user base
- Quickly identify high-performing or struggling users
- Useful for auditing results or providing targeted feedback

This admin console streamlines user management, giving privileged users the power to oversee and support every learner on the platform.

<img width="1785" alt="Screenshot 2025-04-06 at 09 57 30" src="https://github.com/user-attachments/assets/797fa702-fd7d-4e03-8c27-1f433ec3bae2" />


## 🌐 Frontend

The **TestifyApp** frontend is designed for clarity, responsiveness, and ease of maintenance.

### ✨ Features

- **Go HTML Templates**  
  - Pages are rendered server-side using `.gohtml` files  
  - Simplifies dynamic data integration with minimal overhead

- **Tailwind CSS**  
  - Utility-first approach for rapid UI development  
  - Consistent, modern styling across all pages  
  - Easy to extend or override for custom design requirements

- **Plain JavaScript**  
  - Handles interactive elements (timers, validation, button loading states)  
  - Minimal dependencies for better performance and maintainability

- **Responsive Navigation Bar**  
  - Provides quick links to **Home**, **Your Results**, and **All Users** (admin only)  
  - Adaptable design for both desktop and mobile layouts

By combining **Tailwind CSS**, **Go HTML Templates**, and a touch of JavaScript, the frontend remains lightweight, visually consistent, and user-focused.

## 📦 Tech Stack

A high-level overview of the core technologies powering **TestifyApp**:

| Layer          | Technology                                     |
|----------------|------------------------------------------------|
| **Backend**    | [Go](https://go.dev/) + [Fiber](https://gofiber.io/) |
| **Frontend**   | Go HTML templates, Tailwind CSS, Plain JavaScript |
| **Database**   | [AWS RDS MySQL](https://aws.amazon.com/rds/mysql/) + [GORM](https://gorm.io/) |
| **Auth**       | JWT for session tokens + [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for secure password hashing |
| **Deployment** | [AWS EC2](https://aws.amazon.com/ec2/) with a secure TLS connection to AWS RDS |

This stack ensures reliable performance, strong security, and seamless scalability for both development and production environments.


## 📄 .env Example

Below is a sample **`.env`** file showcasing common environment variables and how to configure them. Adjust values according to your application needs:

```bash
DRIVER_NAME=mysql
DB_USER=admin
DB_PASSWORD=yoursecurepassword
DB_HOST=your-rds-endpoint.amazonaws.com
DB_PORT=3306
DB_NAME=yourdbname
DB_CHARSET=utf8mb4
DB_PARSE_TIME=true
DB_TLS_MODE=verify-full
DB_TLS_CERT_PATH=certs/rds-ca-global-bundle.pem

DSN=${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=${DB_CHARSET}&parseTime=${DB_PARSE_TIME}

JWT_SECRET=yourjwtsecret
```

```bash
- **DRIVER_NAME**: Specifies which driver to use (mysql here).  
- **DB_USER** / **DB_PASSWORD**: Credentials for connecting to your RDS instance.  
- **DB_HOST** / **DB_PORT**: RDS endpoint and port.  
- **DB_NAME**: Name of your MySQL database.  
- **DB_CHARSET**: Character set configuration (utf8mb4 recommended).  
- **DB_PARSE_TIME**: Enables time parsing in MySQL responses.  
- **DB_TLS_MODE**: Controls TLS verification (verify-full or verify-ca).  
- **DB_TLS_CERT_PATH**: Path to the SSL certificate for your RDS connection.  
- **DSN**: Constructed from the above variables for easy reference by GORM.  
- **JWT_SECRET**: Secret used for signing JWT tokens.
```

## 🚢 Deployment to AWS

Below is a concise guide for running **TestifyApp** on an **AWS EC2** instance with a secure connection to **AWS RDS MySQL**.

### 1. Prepare the AWS Environment

1. **AWS RDS**  
   - Create a MySQL instance on AWS RDS.  
   - Configure two Security Groups:  
     - **WebTestifySG**: Allows inbound HTTP (port 80) from anywhere.  
     - **DBTestifySG**: Allows inbound MySQL traffic (port 3306) only from WebTestifySG.

2. **AWS EC2**  
   - Launch an EC2 instance (Amazon Linux or Ubuntu).  
   - Associate it with **WebTestifySG**.  
   - Set up or upload an EC2 key pair and download the **.pem** file.

### 2. Build the Go Application

Compile your Go code for Linux by running:

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap
```

This will produce a Linux-compatible binary named **bootstrap**.

### 3. Transfer Files to EC2

Securely copy the compiled binary, TLS certificate, and **.env** file to your EC2 instance.

1)
   ```bash
   scp -i ~/.ssh/WebTestifyAppKP.pem bootstrap ec2-user@<EC2_PUBLIC_IP>:
   ```
   Copies the **bootstrap** file (your Go binary).

2)
   ```bash
   scp -i ~/.ssh/WebTestifyAppKP.pem certs/rds-ca-global-bundle.pem ec2-user@<EC2_PUBLIC_IP>:
   ```
   Transfers the RDS CA certificate for TLS.
3)
   ```bash
   scp -i ~/.ssh/WebTestifyAppKP.pem .env ec2-user@<EC2_PUBLIC_IP>:
   ```
   Uploads your environment variables.

Replace **~/.ssh/WebTestifyAppKP.pem** with the path to your actual key pair, and **<EC2_PUBLIC_IP>** with your instance’s public IP or DNS.

### 4. Launch the Application

After transferring your files:

1) SSH into the EC2 instance (if not already):
   ```bash
   chmod 400 ~/.ssh/WebTestifyAppKP.pem
   ```
   ```bash
   ssh -i ~/.ssh/WebTestifyAppKP.pem ec2-user@<EC2_PUBLIC_IP>
   ```

2) Navigate to the directory containing your files (e.g., **cd ~/** if placed in the home directory).

3) Make the binary executable:
   ```bash
   chmod +x bootstrap
   ```

4) Run the application (with sudo if needed):
   ```bash
   sudo ./bootstrap
   ```

Access your app at **http://<EC2_PUBLIC_IP>** to confirm it is running.
