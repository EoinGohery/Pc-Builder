package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class GpuDto {
    public long Id;
    public String Manufacturer;
    public String GpuName;
    public String Memory;
    public String Clock;
    public Integer Tdp;
    public Integer Price;
}
