
export interface Istudent {
  name: string;
  studentID: number;
  bDate: string;
  position?: string; // ? คือ สามารไม่มีหรือมีก็ได้
  co: Coco[];
}
interface Coco {
  name: string;
}