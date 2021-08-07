import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';

import { ApiReaderService } from './api-reader.service';



describe('ApiReaderService', () => {
  let service: ApiReaderService;

  beforeEach(() => {
    TestBed.configureTestingModule({imports: [HttpClientTestingModule]});
    service = TestBed.inject(ApiReaderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
