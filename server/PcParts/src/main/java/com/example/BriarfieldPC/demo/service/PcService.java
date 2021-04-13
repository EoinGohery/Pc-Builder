package com.example.BriarfieldPC.demo.service;

import com.example.BriarfieldPC.demo.entity.*;
import com.example.BriarfieldPC.demo.repository.*;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@AllArgsConstructor
public class PcService {
    private final CpuRepositroy cpuRepository;
    private final GpuRepository gpuRepository;
    private final DriveRepository driveRepository;
    private final MotherboardRepository motherboardRepository;
    private final PsuRepository psuRepository;
    private final RamRepository ramRepository;

    public List<Cpu> getAllCpu(){
        return cpuRepository.findAll();
    }
    public List<Gpu> getAllGpu(){
        return gpuRepository.findAll();
    }
    public List<Drive> getAllDrivers(){
        return driveRepository.findAll();
    }
    public List<Motherboard> getAllMotherboards(){
        return motherboardRepository.findAll();
    }
    public List<PSU> getAllPSU(){
        return psuRepository.findAll();
    }
    public List<Ram> getAllRam(){
        return ramRepository.findAll();
    }
}
