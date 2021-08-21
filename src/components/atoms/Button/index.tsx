import React, { FC, MouseEventHandler } from 'react';
import styles from './styles.module.css';

interface Props {
  children: React.ReactNode;
  className?: string;
  onClick?: MouseEventHandler<HTMLImageElement>;
}

const Button: FC<Props> = ({ children, className}) => {
  return (
    <button className={[styles.button, className].join(' ')}>{children}</button>
  )
}

export default Button;