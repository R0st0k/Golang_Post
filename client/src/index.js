import React from 'react';
import ReactDOM from 'react-dom/client';
import {
    createBrowserRouter,
    RouterProvider,
    Route,
} from "react-router-dom";
import './styles/index.css';
import ErrorPage from "./pages/error-page.jsx";
import OrderTracking from "./pages/order-tracking"
import Main from "./pages/main"
import Registration from "./pages/registration";

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
    }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <RouterProvider router={router} />
  </React.StrictMode>
);

