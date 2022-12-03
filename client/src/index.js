import React from 'react';
import ReactDOM from 'react-dom/client';
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import './styles/index.css';
import ErrorPage from "./pages/error-page.jsx";
import OrderTracking from "./pages/order-tracking"
import Main from "./pages/main"
import Registration from "./pages/registration";
import SendingsInfo from "./pages/sendings-info";
import Employees from "./pages/employees";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Main />,
        errorElement: <ErrorPage />,
    },
    {
        path: "orders/:orderId",
        element: <OrderTracking />
    },
    {
        path: "/registration",
        element: <Registration />
    },
    {
        path: "/sendings-info",
        element: <SendingsInfo />
    },
    {
        path: "/employees",
        element: <Employees/>
    }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <RouterProvider router={router} />
  </React.StrictMode>
);

