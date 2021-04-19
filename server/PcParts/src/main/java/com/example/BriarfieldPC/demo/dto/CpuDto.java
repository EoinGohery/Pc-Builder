package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CpuDto {
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Cores;
    public String Clock;
    public Integer Tdp;
    public String Socket;
    public Integer Price;
}
