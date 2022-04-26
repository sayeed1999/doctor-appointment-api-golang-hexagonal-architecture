import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiResponse } from '../../models/api-response.model';

@Injectable({
  providedIn: 'root'
})
export class RepositoryService {

  protected url = 'http://localhost:8080'
  protected postUrl = '' // added url to the last

  constructor(protected http: HttpClient) { }

  getAll(): Observable<ApiResponse> {
    return this.http.get<ApiResponse>(`${this.url}/${this.postUrl}`)
  }

  getById(id: number): Observable<ApiResponse> {
    if(this.postUrl.length > 0)
      return this.http.get<ApiResponse>(`${this.url}/${this.postUrl}/${id}`)
    
    return this.http.get<ApiResponse>(`${this.url}/${id}`)
  }

  post(obj: object): Observable<ApiResponse> {
    return this.http.post<ApiResponse>(
      this.url + '/' + this.postUrl,
      obj
    )
  }

  put(obj: object): Observable<ApiResponse> {
    return this.http.put<ApiResponse>(
      this.url + '/' + this.postUrl,
      obj
    )
  }

  delete(obj: object): Observable<ApiResponse> {
    return this.http.delete<ApiResponse>(
      this.url + '/' + this.postUrl,
      obj
    ) 
  }
}
