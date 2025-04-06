# 🧪 TestifyApp

**TestifyApp** is a full-stack web application built with Go and Fiber, designed for creating, taking, and managing tests with role-based access for users and admins. It features JWT authentication, dynamic test-taking functionality, result tracking, and secure deployment on AWS.

---

## 🚀 Features

- ✅ **User Authentication & Authorization**
    - Sign up and login using JWT tokens
    - Passwords securely hashed with `bcrypt`
    - JWT stored in cookies for secure session management


- 🏁 **Test System**
    - Timed tests with a combination of single and multi-choice questions
    - Immediate result evaluation and display upon test submission
    - Score tracking and pass/fail feedback


- 📊 **User Dashboard**
    - View personal test history and results
    - Score, duration, attempt count, and pass status shown in a sortable table


- 🔐 **Admin Dashboard**
    - Flag users as `is_admin = true` in the DB to grant admin access
    - View all registered users
    - Inspect individual user results and attempts


- 🌐 **Frontend**
    - Rendered using `gohtml` templates
    - Dynamic behavior powered by plain JavaScript (form validation, button loading states, timer, etc.)
    - Responsive navbar with navigation between Home, Results, and Admin pages


- 📦 **Tech Stack**
    - **Backend:** Golang + [Fiber](https://github.com/gofiber/fiber)
    - **Frontend:** HTML (gohtml), CSS, JavaScript
    - **Database:** AWS RDS MySQL using GORM ORM
    - **Authentication:** JWT + Bcrypt
    - **Hosting:** AWS EC2 + Secure TLS Connection to RDS

---

## 🛠️ Technologies Used

| Layer        | Technology                         |
|--------------|-------------------------------------|
| Backend      | Go + Fiber                          |
| ORM          | GORM                                |
| Frontend     | gohtml templates + JS               |
| Auth         | JWT + bcrypt                        |
| Database     | AWS RDS MySQL                       |
| Deployment   | AWS EC2 (with TLS to RDS)           |

---

## 📸 Screenshots

> _(Include screenshots or a GIF demo if possible to visually showcase the app flow — sign up, login, take test, see results, admin panel, etc.)_

---

## 🧑‍💻 Local Development Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/testifyapp.git
cd testifyapp