import React, { FC } from 'react';
import './styles.css';

interface Props {
  children: React.ReactNode;
  className?: string;
}

const Button: FC<Props> = ({ children, className}) => {
  return (
    <button className={['button', className].join(' ')}>{children}</button>
  )
}

export default Button;