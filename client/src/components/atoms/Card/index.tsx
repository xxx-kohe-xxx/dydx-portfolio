import React, { FC } from 'react';
import styles from './styles.module.css';

interface Props {
  children: React.ReactNode;
  className?: string;
}

const Card: FC<Props> = ({ children, className }) => {
  return (
    <section className={[styles.card, className].join(' ')}>
      {children}
    </section>
  )
}

export default Card;

export const CardHeader: FC<Props> = ({ children, className}) => {
  return (
    <div className={[styles.header, className].join(' ')}>
      {children}
    </div>
  )
}