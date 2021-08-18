import React, { FC } from 'react';
import './styles.css';
import TextBox from '../../atoms/TextBox';
import Button from '../../atoms/Button';

interface Props {
  onSubmit: (event: MouseEvent) => void;
}

const InputForm: FC<Props> = ({ onSubmit }) => {
  return (
    <form className='form'>
      <TextBox className='textbox'/>
      <Button onClick={onSubmit}>追加</Button>
    </form>
  )
}

export default InputForm;