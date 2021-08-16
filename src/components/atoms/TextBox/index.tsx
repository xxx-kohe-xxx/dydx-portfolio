import React, { FC } from 'react';
import './styles.css';

interface Props {
  className?: string;
}

const TextBox: FC<Props> = ({ className }) => {
  return (
    <input type="text" className={['textbox', className].join(' ')} />
  )
}

export default TextBox;