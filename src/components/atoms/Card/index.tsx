import React, { FC } from 'react';
import './styles.css';

interface Props {
  children: React.ReactNode;
  className?: string;
}

const Card: FC<Props> = ({ children, className }) => {
  return (
    <section className={['card', className].join(' ')}>
      {children}
    </section>
  )
}

export default Card;

export const CardHeader: FC<Props> = ({ children, className}) => {
  return (
    <div className={['header', className].join(' ')}>
      {children}
    </div>
  )
}