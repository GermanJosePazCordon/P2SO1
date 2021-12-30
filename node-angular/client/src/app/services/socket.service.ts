import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { io } from "socket.io-client";

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  public message$: BehaviorSubject<string> = new BehaviorSubject('');
  constructor() {}

  socket = io('172.17.0.2:4000');

  public getNewMessage = () => {
    this.socket.on('data', (message) =>{
      this.message$.next(message);
    });
    
    return this.message$.asObservable();
  };
}
