import { AxiosRequestConfig } from 'axios';

export const setLocalToken = (token: string) => {
  localStorage.setItem('token', token);
}

export const removeLocalToken = () => {
  localStorage.removeItem('token');
};

export const getLocalToken = () => {
  return localStorage.getItem('token');
};

export const getAuthorizationHeader = (token: string) => {
  return {
    headers: { Authorization: `Bearer ${token}` },
  } as AxiosRequestConfig;
};

export const formatDate = (date: Date) => {
  const dayNumber = date.getDate().toString();
  const month = (date.getMonth() + 1).toString();
  const year = date.getFullYear().toString();

  const hours = date.getHours().toString();
  const minutes = date.getMinutes().toString();

  return `${dayNumber}/${month}/${year} às ${hours}:${minutes}`;
}
