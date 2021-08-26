import React, { FC } from 'react';
import styles from './styles.module.css';
import TextBox from '../../atoms/TextBox';
import Button from '../../atoms/Button';

interface Props {
  onSubmit: React.MouseEventHandler<HTMLImageElement>;
}

const InputForm: FC<Props> = ({ onSubmit }) => {
  return (
    <form className={styles.form}>
      <TextBox className={styles.textbox}/>
      <Button onClick={onSubmit}>追加</Button>
    </form>
  )
}

export default InputForm;