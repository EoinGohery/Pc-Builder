package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class PSUDto {
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Price;
    public Integer Capacity;
    public String Rating;
}
