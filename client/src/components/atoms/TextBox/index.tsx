import React, { FC } from 'react';
import styles from './styles.module.css';

interface Props {
  className?: string;
}

const TextBox: FC<Props> = ({ className }) => {
  return (
    <input type="text" className={[styles.textbox, className].join(' ')} />
  )
}

export default TextBox;