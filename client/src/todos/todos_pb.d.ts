import * as jspb from 'google-protobuf'



export class Todos extends jspb.Message {
  getTodosList(): Array<Todo>;
  setTodosList(value: Array<Todo>): Todos;
  clearTodosList(): Todos;
  addTodos(value?: Todo, index?: number): Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todos.AsObject;
  static toObject(includeInstance: boolean, msg: Todos): Todos.AsObject;
  static serializeBinaryToWriter(message: Todos, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todos;
  static deserializeBinaryFromReader(message: Todos, reader: jspb.BinaryReader): Todos;
}

export namespace Todos {
  export type AsObject = {
    todosList: Array<Todo.AsObject>,
  }
}

export class Todo extends jspb.Message {
  getId(): number;
  setId(value: number): Todo;

  getTitle(): string;
  setTitle(value: string): Todo;

  getDone(): boolean;
  setDone(value: boolean): Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
    id: number,
    title: string,
    done: boolean,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

